import { Page, Locator, expect } from '@playwright/test';

export default class StudentBusTrackPage {
    readonly page: Page;

    // Define URLs
    readonly signInUrl: string = '/profile';
    readonly busTrackUrl: string = '/bus-track';
    readonly emailInput: Locator;
    readonly passwordInput: Locator;
    readonly signInButton: Locator;
    readonly signInPageHeader: Locator;
    readonly busArrivalStatus: Locator;
    readonly stopDetails: Locator;

    constructor(page: Page) {
        this.page = page;
        this.signInPageHeader = page.locator("h1:has-text('Sign in')");
        this.emailInput = page.locator("[data-testid='sign-in-email-input']");
        this.passwordInput = page.locator("[data-testid='sign-in-password-input']");
        this.signInButton = page.locator("button:has-text('Sign in')");
        this.busArrivalStatus = page.locator("[data-testid='bus-arrival-status']");
        this.stopDetails = page.locator("[data-testid='stop-details']");
    }

    // Navigate to Sign-In Page
    async visitSignIn() {
        await this.page.goto(this.signInUrl);
    }

    // Sign In
    async signIn(email: string, password: string) {
        await this.emailInput.fill(email);
        await this.passwordInput.fill(password);
        await this.signInButton.click();
    }

    async visitBusTrack() {
        await this.page.goto(this.busTrackUrl);
    }

    async simulateBusArrival() {

    }

    async checkBusArrivalStatus(expectedText: string) {
        await expect(this.busArrivalStatus).toHaveText(expectedText);
    }

    async checkStopDetailsVisible() {
        await expect(this.stopDetails).toBeVisible();
    }
}
