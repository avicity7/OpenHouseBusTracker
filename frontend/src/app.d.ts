declare global {
	declare namespace App {
		interface Locals {
			session: {user: {username: string, email: string }} | null
		}
	}
}

export {}