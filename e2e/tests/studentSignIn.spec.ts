import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';
import DashboardPage from '../models/DashboardPage';

test.describe('Student User Sign In', () => {
  test('should sign in as a student with valid credentials', async ({ page }) => {
    const signInPage = new SignInPage(page);
    await signInPage.visit();

    const email = 'user@gmail.com';
    const password = '1';
    await signInPage.signIn(email, password);

    // Assuming the student is redirected to a dashboard or specific page
    const dashboardPage = new DashboardPage(page);
    await expect(dashboardPage.studentWelcomeMessage).toBeVisible();
   
  });

  test('should display error message for student with invalid credentials', async ({ page }) => {
    const signInPage = new SignInPage(page);
    await signInPage.visit();

    const email = 'invalid_student@example.com';
    const password = 'wrong_password';
    await signInPage.signIn(email, password);


  });
});
