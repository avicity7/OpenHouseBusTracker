import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { SwapRequest } from "$lib/types/global";

export const load = async () => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-accepted-swaps`);
  if (!response.ok) {
      throw new Error("Failed to fetch event helpers");
  }
  
  let swaps = await response.json() as SwapRequest[];
  if (!swaps) {
    swaps = []
  }

  return {
    swaps
  }
}