import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('View Schedule', () => {
    test('Available Drivers', async ({ page }) => {
        // Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'sarahaxl7@gmail.com';
        const password = '123';
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/schedule')) {
            console.log('Navigating manually to /admin/schedule');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/schedule');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/schedule');
        await page.screenshot({ path: 'admin-schedules-page.png' });

        // Add Schedule
        const addScheduleButton = await page.locator('[data-testid="add-schedule-button"]');
        await addScheduleButton.click();
        await page.goto('https://open-house-bus-tracker.vercel.app/admin/schedule/add-schedule');

        // Driver
        const driverDropdown = await page.locator('[data-testid="driver_id"]');
        await driverDropdown.click(); 
        await page.waitForSelector('ul[role="listbox"] li button');
        const availableDrivers = await page.locator('ul[role="listbox"] li button').allTextContents();
        console.log('Available drivers in the dropdown:', availableDrivers);
        expect(availableDrivers.length).toBeGreaterThan(0);
        for (const driver of availableDrivers) {
            console.log('Available driver:', driver);
        }
        await page.screenshot({ path: 'available-drivers-dropdown.png' });
        await driverDropdown.press('Escape');
    });
});
