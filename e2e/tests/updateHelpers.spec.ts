import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Reassign Bus and Drivers', () => {
    test('Update Student Helpers', async ({ page }) => {
        //Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'sarahaxl7@gmail.com'; 
        const password = '123'; 
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/event-helper')) {
            console.log('Navigating manually to /admin/event-helper');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/event-helper');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/event-helper');
        await page.screenshot({ path: 'admin-event-helper-page.png' });

        //Update Helper Page
        const addScheduleButton = await page.locator('[data-testid="sarah"]');
        await addScheduleButton.click();
        await page.goto('https://open-house-bus-tracker.vercel.app/admin/event-helper/update-helper');

        // Select Carplate
        const carPlateDropdown = await page.locator("[data-testid='carplate']");
        await carPlateDropdown.click(); 
        await page.waitForSelector('ul[role="listbox"] li button');
        await page.locator('ul[role="listbox"] li button:has-text("Yellow 5")').click(); 

        //Select Shift
        const shiftRadio = await page.locator("[data-testid='false']");
        await shiftRadio.click();
        await expect(shiftRadio).toBeChecked();

        // Submit Form
        const submitButton = await page.locator('button[type="submit"]');
        await submitButton.click();
        console.log('Update helpers successfully.');

    });
});
