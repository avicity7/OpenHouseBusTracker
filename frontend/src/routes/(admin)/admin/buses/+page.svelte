<script lang="ts">
  export let data
  let { buses, routes, scheduleBus, backend_uri, env } = data
  import type { EventBus, Demand, BusAssignments } from '$lib/types/global.js';
	import { onMount } from 'svelte';
  import ToolTip from '$lib/components/ToolTip.svelte';
	import ToggleSwitch from '$lib/components/ToggleSwitch.svelte';
	import StateMenuBar from '$lib/components/StateMenuBar.svelte';
	let ws: WebSocket;

  let showHidden = false
  let autoScheduling = false
   type RouteStates = {
    [key: string]: string;
  };

  let routeStates: RouteStates = {};
  
  const getBuses = async() => {
    const response = await fetch(`${backend_uri}:3000/bus/get-buses`)
    if (response) {
      buses = await response.json() as Array<EventBus> 
    }
  }

  const updateBusVisibility = async(carplate: string, visibility: boolean) => {
    await fetch(`${backend_uri}:3000/bus/update-bus-visibility/${carplate}/${visibility}`, {
      method: 'PUT'
    })
    getBuses()
  }
    
  const deleteBus = async(carplate: string) => {
    await fetch(`${backend_uri}:3000/bus/delete-bus/${carplate}`, {
      method: 'DELETE'
    })
    getBuses()
  }


  const updateBusState = async (routeName: string, newState: string) => {
    console.log(`Route: ${routeName}, Current State: ${newState}`);
    routeStates[routeName] = newState;
    if (autoScheduling) {
      autoScheduleBuses();
    }
  };

  const toggleAutoScheduling = async (newChecked: boolean) => {
    autoScheduling = newChecked;
    if (autoScheduling) {
      autoScheduleBuses();
    }
  };

  const autoScheduleBuses = () => {
    let demand: Demand = {};
    for (let route of routes) {
      let state = routeStates[route.RouteName];
      demand[route.RouteName] = state === 'Full Bus' ? 3 : state === 'Half Full' ? 2 : 1;
    }

    let touringBuses = buses.filter(bus => bus.Status);
    touringBuses = touringBuses.filter(bus => scheduleBus.some(sb => sb.Carplate === bus.Carplate));

    let initialAssignments = scheduleBus.filter(sb => touringBuses.some(tb => tb.Carplate === sb.Carplate));

    let totalBuses = touringBuses.length;
    console.log("intitial ", initialAssignments)
    console.log("touring buses", touringBuses);

    // calculation of importance for buses, could change back to fix assignment of buses if needed
    // some issue with pop and push behaviour, to be clarified
    let busAssignments: BusAssignments = {};
    let totalDemand = Object.values(demand).reduce((a, b) => a + b, 0);
    for (let route in demand) {
      busAssignments[route] = [];
      let numBusesForRoute = Math.round((demand[route] / totalDemand) * totalBuses);
      for (let i = 0; i < numBusesForRoute && touringBuses.length > 0; i++) {
        let bus = touringBuses.pop();
        if (bus) {
          busAssignments[route].push(bus.Carplate);
        }
      }
    }

    console.log('Bus Assignments:', busAssignments);

    let assignmentData = [];
    for (let route in busAssignments) {
      for (let carplate of busAssignments[route]) {
        let initialAssignment = initialAssignments.find(ia => ia.Carplate === carplate);
        if (initialAssignment && initialAssignment.RouteName !== route) {
          assignmentData.push({ Carplate: carplate, RouteName: route });
        }
      }
    }

    console.log('Assignment Data:', assignmentData);

    const assignmentDataField = document.getElementById('assignmentData') as HTMLInputElement;
      assignmentDataField.value = JSON.stringify(assignmentData);

    const form = document.getElementById('assignmentForm') as HTMLFormElement;
    form.submit();
  };

  onMount(() => {
		ws = new WebSocket(`${env == 'PROD' ? 'wss' : 'ws'}://${backend_uri.split('//')[1]}:3000/ws`);

    ws.onmessage = async (msg) => {
      let parts = msg.data.split(' ');
      if (parts[0] == 'refresh') {
        getBuses()
      }
    }
  })
</script>

<form id="assignmentForm" method="POST" action="?/updateScheduleRoutes">
  <input type="hidden" id="assignmentData" name="assignmentData" />
</form>

<div class="p-6 md:p-12">
  <div class="flex justify-between">
    <div class="flex items-end">
      <div>
        <a href="/admin/buses/create-bus" class="w-fit bg-red-700 hover:bg-red-800 px-8 py-2 my-6 text-white font-semibold rounded-md">
          Add Bus
        </a>
      </div>
      <div>
        <input class="ml-4" type="checkbox" id="hidden" name="hidden" checked={showHidden} on:change={() => { showHidden = !showHidden }}>
        <label for="hidden">Show hidden buses</label>
      </div>
    </div>
  
    <div class="bg-red-100 p-4 rounded-lg shadow-md w-full max-w-md ml-4">
      <div class="flex items-center justify-between mb-4">
        <span class="text-md font-semibold">Auto Scheduling</span>
        <ToggleSwitch checked={autoScheduling} onToggle={toggleAutoScheduling} />
      </div>
      
      {#each routes as route}
        <div class="flex items-center justify-between mb-2">
          <h1 class="mr-4 text-sm">{route.RouteName}</h1>
          <div class="w-full">
            <StateMenuBar 
              initialState={""}
              onStateChange={(newState) => updateBusState(route.RouteName, newState)} 
            />
          </div>
        </div>
      {/each}
    </div>
  </div>
  

  <div class="mt-8">
    <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
            <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Carplate</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                <!-- <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">State</th> -->
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Action</th>
            </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
                {#each buses as bus}
                  {#if !bus.Hidden || showHidden}
                    <tr class="hover:bg-gray-100">
                        <td class="px-6 py-4 whitespace-nowrap">{bus.Carplate}</td>
                        <td class="px-4 py-6">
                          <p class={bus.Status ? "text-green-600" : "text-orange-600"}>{bus.Status ? 'Touring' : 'Inactive'}</p>
                        </td>
                        <!-- <td class="px-6 py-4 whitespace-nowrap">
                          <StateMenuBar state={"true"} onStateChange={(newState) => updateBusState(bus.Carplate, newState)} />
                        </td> -->
                    
                        <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                            <div class="flex items-center">
                                <button class="text-stone-500 hover:text-red-600 text-2xl mr-4" on:click={() => updateBusVisibility(bus.Carplate, !bus.Hidden)}>
                                    <ToolTip text={bus.Hidden ? "Show Bus" : "Hide Bus"}> 
                                      {#if bus.Hidden}
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16m-3.5-3.244C15.147 17.485 13.618 18 12 18c-3.53 0-6.634-2.452-8.413-4.221c-.47-.467-.705-.7-.854-1.159c-.107-.327-.107-.913 0-1.24c.15-.459.385-.693.855-1.16c.897-.892 2.13-1.956 3.584-2.793M19.5 14.634c.333-.293.638-.582.912-.854l.003-.003c.468-.466.703-.7.852-1.156c.107-.327.107-.914 0-1.241c-.15-.458-.384-.692-.854-1.159C18.633 8.452 15.531 6 12 6c-.338 0-.671.022-1 .064m2.323 7.436a2 2 0 0 1-2.762-2.889"/></svg>
                                      {:else}
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.587 13.779c1.78 1.769 4.883 4.22 8.413 4.22c3.53 0 6.634-2.451 8.413-4.22c.47-.467.705-.7.854-1.159c.107-.327.107-.913 0-1.24c-.15-.458-.385-.692-.854-1.159C18.633 8.452 15.531 6 12 6c-3.53 0-6.634 2.452-8.413 4.221c-.47.467-.705.7-.854 1.159c-.107.327-.107.913 0 1.24c.15.458.384.692.854 1.159"/><path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/></g></svg>
                                      {/if}
                                    </ToolTip>
                                </button>
                                <a href={`buses/update-bus/${encodeURIComponent(JSON.stringify(bus))}`} class="text-stone-500 hover:text-green-500 mr-8">
                                    <ToolTip text="Update Bus"> 
                                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                        <path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/>
                                      </svg>
                                    </ToolTip>
                                </a>
                                <button class="text-stone-500 hover:text-red-600 text-2xl" on:click={() => deleteBus(bus.Carplate)}>
                                    <ToolTip text="Delete Bus"> 
                                    <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                            <g fill="none">
                                                <path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
                                                <path fill="currentColor" d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z" />
                                            </g>
                                        </svg>
                                    </svg>
                                    </ToolTip>
                                </button>
                            </div>
                        </td>                 
                    </tr>
                  {/if}
                {/each}
        </tbody>
    </table>
  </div>
</div>


<!-- <div class="max-w-sm md:max-w-4xl mx-auto bg-white p-2 rounded-lg">
    <table class="w-full">
      <thead>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Carplate</th>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Modify</th>
      </thead>
      <tbody class="divide-y divide-stone-300">
        {#each buses as bus}
          <tr>
            <td class="px-4 py-6">{bus.Carplate}</td>
            <td class="px-4 py-6">
              <p class={bus.Status ? "text-green-600" : "text-orange-600"}>{bus.Status ? 'Touring' : 'Inactive'}</p>
            </td>
            <td class="px-4 py-6">
              <button class="text-stone-500 hover:text-red-600 text-2xl" on:click={() => deleteBus(bus.Carplate)}>
                  <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                          <g fill="none">
                              <path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
                              <path fill="currentColor" d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z" />
                          </g>
                      </svg>
                  </svg>
              </button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div> -->