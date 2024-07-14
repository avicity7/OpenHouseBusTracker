import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Add Schedule', () => {
    test('Add Schedule to assign student helpers', async ({ page }) => {
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

        // Select Carplate
        const carPlateDropdown = await page.locator("[data-testid='carplate']");
        await carPlateDropdown.selectOption({ label: 'Green 2' });

        // Select Route Name
        await page.locator('label:has-text("Route Name") + div .dropdown-toggle').click();
        await page.locator('label:has-text("Route Name") + div ul[role="listbox"] li button:has-text("Green Route")').click();

        // Select Driver
        await page.locator('label:has-text("Driver") + div .dropdown-toggle').click();
        await page.locator('label:has-text("Driver") + div ul[role="listbox"] li button:has-text("John Smith")').click();

        // Fill Start Time
        const startTimeInput = await page.locator('[data-testid="start-time"]');
        await startTimeInput.fill('2024-06-01T08:30');

        // Fill End Time
        const endTimeInput = await page.locator('[data-testid="start-time"]');
        await endTimeInput.fill('2024-06-01T17:30');

        // Submit Form
        const submitButton = await page.locator('button[type="submit"]');
        await submitButton.click();
        console.log('Schedule added successfully.');
    });
});
