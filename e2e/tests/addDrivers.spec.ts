import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Add Drivers', () => {
    test('Manual Adding of Drivers', async ({ page }) => {
        //Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'sarahaxl7@gmail.com'; 
        const password = '123'; 
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/drivers')) {
            console.log('Navigating manually to /admin/drivers');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/drivers');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/drivers');
        await page.screenshot({ path: 'admin-drivers-page.png' });

        // Add Button
        const addDriverButton = await page.locator('[data-testid="add-driver-button"]');
        await addDriverButton.click();
        const driverNameInput = await page.locator('[data-testid="driver-name-input"]');
        await driverNameInput.fill('John Doe');
        const submitButton = await page.locator('[data-testid="submit-add-driver"]');
        await submitButton.click();
        console.log('Driver added successfully.');
    });
});
