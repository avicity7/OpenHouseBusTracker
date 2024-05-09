import { test } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Sign In Page', () => {
  test('should sign in with valid credentials', async ({ page }) => {
    const signInPage = new SignInPage(page);
    await signInPage.visit();

    const email = 'user@gmail.com';
    const password = '1';
    await signInPage.signIn(email, password);
  });


  // no error message done in FE yet
  test('should display error message with invalid credentials', async ({ page }) => {
    const signInPage = new SignInPage(page);
    await signInPage.visit();

    const email = 'invalid_email@example.com';
    const password = 'invalid_password';
    await signInPage.signIn(email, password);
  });
});