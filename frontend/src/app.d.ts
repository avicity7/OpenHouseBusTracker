declare global {
	declare namespace App {
		interface Locals {
			session: { Name: string, Email: string, Role: string, VerificationToken: string } | null,
		}
	}
}

export {}