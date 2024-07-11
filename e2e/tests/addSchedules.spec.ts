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
        const carplateDropdown = await page.locator('select[name="carplate"]');
        await carplateDropdown.selectOption({ label: 'Green 2' }); 
        const routeNameDropdown = await page.locator('select[name="route_name"]');
        await routeNameDropdown.selectOption({ label: 'Green Route' }); 
        const driverDropdown = await page.locator('select[name="driver_id"]');
        await driverDropdown.selectOption({ label: 'John Smith' }); 
        const startTimeInput = await page.locator('input[name="start_time"]');
        await startTimeInput.fill('01/06/2024 08.30 am');
        const endTimeInput = await page.locator('input[name="end_time"]');
        await endTimeInput.fill('01/06/2024 05.30 pm');
        const submitButton = await page.locator('button[type="submit"]');
        await submitButton.click();
        console.log('Schedule added successfully.');
    });
});
