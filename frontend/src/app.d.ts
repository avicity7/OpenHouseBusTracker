declare global {
	declare namespace App {
		interface Locals {
			session: { Email: string, Role: string } | null,
		}
	}
}

export {}