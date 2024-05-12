<script lang="ts">
    export let data
    const { backend_uri } = data
    
    let routing: string
  
    const createRoute = async() => {
      await fetch(`${backend_uri}:3000/route/create-route`, {
        method: 'POST',
        body: JSON.stringify({
            RouteName: routing
        })
      })
      location.replace('/admin/bus-routes')
    }
  </script>
  <div class="flex justify-center items-center h-full">
      <div class="bg-white shadow rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-12">
          <h1 class="text-2xl font-semibold mb-4">Add New Route</h1>
          <form on:submit|preventDefault={createRoute}>
              <div class="mb-4">
                  <label for="AddiingRoute" class="block text-sm font-medium mb-1">Route Name</label>
                  <input
                      type="text"
                      required
            bind:value={routing}
                      class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
                  />
              </div>
  
              <div class="mt-4 flex justify-center">
                  <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
                      >Add Route</button
                  >
              </div>
          </form>
      </div>
  </div>
  