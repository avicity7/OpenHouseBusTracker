<svelte:head>
	<title>AI Chatbot Help | SPOH Bus Tracker</title>
	<style>
		#webchat {
			height: calc(100% - 50px);
			overflow: hidden;
			position: fixed;
			top: 50px;
			width: 100%;
		}
	</style>
	<script
		bind:this={script}
		crossorigin="anonymous"
		src="https://cdn.botframework.com/botframework-webchat/latest/webchat.js"
	></script>
</svelte:head>

<script lang="ts">
	import { onMount } from "svelte";

	let script: HTMLElement;
	onMount(async function () {
		script.addEventListener("load", async () => {
			const styleOptions = {
				hideUploadButton: true
			};

			const tokenEndpointURL = new URL(
				'https://default2dab15094fb44901ab2d209deb2f26.db.environment.api.powerplatform.com/powervirtualagents/botsbyschema/cr352_athena/directline/token?api-version=2022-03-01-preview'
			);

			const locale = document.documentElement.lang || 'en'; // Uses language specified in <html> element and fallback to English (United States).

			const apiVersion = tokenEndpointURL.searchParams.get('api-version');

			const [directLineURL, token] = await Promise.all([
				fetch(
					new URL(
						`/powervirtualagents/regionalchannelsettings?api-version=${apiVersion}`,
						tokenEndpointURL
					)
				)
					.then((response) => {
						if (!response.ok) {
							throw new Error('Failed to retrieve regional channel settings.');
						}

						return response.json();
					})
					.then(({ channelUrlsById: { directline } }) => directline),
				fetch(tokenEndpointURL)
					.then((response) => {
						if (!response.ok) {
							throw new Error('Failed to retrieve Direct Line token.');
						}

						return response.json();
					})
					.then(({ token }) => token)
			]);

			const directLine = WebChat.createDirectLine({
				domain: new URL('v3/directline', directLineURL),
				token
			});

			const subscription = directLine.connectionStatus$.subscribe({
				next(value) {
					if (value === 2) {
						directLine
							.postActivity({
								localTimezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
								locale,
								name: 'startConversation',
								type: 'event'
							})
							.subscribe();

						subscription.unsubscribe();
					}
				}
			});

			WebChat.renderWebChat(
				{ directLine, locale, styleOptions },
				document.getElementById('webchat')
			);
		})
	});
</script>

<div class="p-8" id="webchat" role="main"></div>