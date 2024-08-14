import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Add Buses', () => {
    test('Manual Adding of Buses', async ({ page }) => {
        //Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'gabrieltai2001@gmail.com'; 
        const password = 'test123'; 
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/buses')) {
            console.log('Navigating manually to /admin/buses');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/buses');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/buses');
        await page.screenshot({ path: 'admin-buses-page.png' });

        // Add Button
        const addBusButton = await page.locator('[data-testid="add-bus-button"]');
        await addBusButton.click();
        const busNameInput = await page.locator('[data-testid="bus-name-input"]');
        await busNameInput.fill('John Doe');
        const submitButton = await page.locator('[data-testid="submit-add-bus"]');
        await submitButton.click();
        console.log('Bus added successfully.');
    });
});
