import { Page, Locator, expect } from "@playwright/test";

export default class SignInPage {
    readonly page: Page;
  
    readonly url: string = '/profile';
  
    readonly emailInput: Locator;
    readonly passwordInput: Locator;
    readonly signInButton: Locator;
    readonly signInPageHeader: Locator;
  
    constructor(page: Page) {
        this.page = page;

        this.signInPageHeader = page.locator("h1:has-text('Sign in')");
  
        this.emailInput = page.locator("[data-testid='sign-in-email-input']");
        this.passwordInput = page.locator("[data-testid='sign-in-password-input']");
        this.signInButton = page.locator("button:has-text('Sign in')");
    }
  
    async visit() {
        await this.page.goto(this.url);
    }

    async isSignInHeaderVisible() {
        await expect(this.signInPageHeader).toBeVisible();
    }
  
    async signIn(email: string, password: string) {
        await this.emailInput.fill(email);
        await this.passwordInput.fill(password);
        await this.signInButton.click();
    }
}
