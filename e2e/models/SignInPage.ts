import { type Page, type Locator, expect } from "@playwright/test";

export default class SignInPage {
    readonly page: Page;
  
    readonly url: string = 'http://localhost:5173/profile';
  
    readonly emailInput: Locator;
    readonly passwordInput: Locator;
    readonly signInButton: Locator;
    readonly signInPageHeader: Locator;
  
    constructor(page: Page) {
      this.page = page;

      this.signInPageHeader = page.getByRole("heading", {
        name: "Sign In",
    });
  
      this.emailInput = page.getByTestId("sign-in-email-input");
      this.passwordInput = page.getByTestId("sign-in-password-input");
      this.signInButton = page.getByRole('button', {
            name: "Sign In",
        });

    }
  
    async visit() {
      await this.page.goto(this.url);
    }

    async isSignInHeaderVisible() {
        await expect(this.signInPageHeader).toBeVisible();
    }
  
    async signIn(email: string, password: string) {
        await this.emailInput.fill(email)
        await this.passwordInput.fill(password);
        await this.signInButton.click();;
      }
  }


//   data-testid="no-announcements"

