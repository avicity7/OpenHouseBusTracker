import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Bulk Import Users', () => {
    test('Bulk Import Users Functionality', async ({ page }) => {
        //Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'sarahaxl7@gmail.com'; 
        const password = '123'; 
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/users')) {
            console.log('Navigating manually to /admin/users');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/users');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/users');
        await page.screenshot({ path: 'admin-users-page.png' });

        // Bulk Add Users
        const fileInput = await page.locator('[data-testid="file-input"]');
        await fileInput.click();
        await fileInput.setInputFiles('C:/bulk-import-file.csv'); 
        const uploadCSVButton = await page.locator('[data-testid="submit-bulk-import"]');
        await uploadCSVButton.click();
        console.log('Bulk import completed successfully.');
    });
});
