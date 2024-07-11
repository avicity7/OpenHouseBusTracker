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

        // Select Carplate
        const carplateDropdown = await page.locator('custom-dropdown[label="Carplate"] select');
        await carplateDropdown.click();
        await page.waitForSelector('custom-dropdown[label="Carplate"] select option[value="Green 2"]');
        await carplateDropdown.selectOption({ value: 'Green 2' });

        // Select Route Name
        const routeNameDropdown = await page.locator('custom-dropdown[label="Route Name"] select');
        await routeNameDropdown.click();
        await page.waitForSelector('custom-dropdown[label="Route Name"] select option[value="Green Route"]');
        await routeNameDropdown.selectOption({ value: 'Green Route' });

        // Select Driver
        const driverDropdown = await page.locator('custom-dropdown[label="Driver"] select');
        await driverDropdown.click();
        await page.waitForSelector('custom-dropdown[label="Driver"] select option[value="John Smith"]');
        await driverDropdown.selectOption({ value: 'John Smith' });

        // Fill Start Time
        const startTimeInput = await page.locator('input[name="start_time"]');
        await startTimeInput.fill('2024-06-01T08:30');

        // Fill End Time
        const endTimeInput = await page.locator('input[name="end_time"]');
        await endTimeInput.fill('2024-06-01T17:30');

        // Submit Form
        const submitButton = await page.locator('button[type="submit"]');
        await submitButton.click();
        console.log('Schedule added successfully.');
    });
});
