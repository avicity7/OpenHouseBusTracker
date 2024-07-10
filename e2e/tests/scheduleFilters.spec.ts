import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Schedule Filters', () => {
    test('Filter Schedules by Route, Car Plate, Start Time and End Time', async ({ page }) => {
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
        await page.screenshot({ path: 'admin-schedule-page.png' });

        // Route
        const routeDropdown = await page.locator("[data-testid='search-route']");
        console.log('Route dropdown locator:', routeDropdown);
        await routeDropdown.selectOption({ label: 'Green Route' });
        console.log('Selected "Green Route" from route dropdown');
        await page.waitForTimeout(2000);

        const scheduleItems = await page.$$('.schedule-item');
        console.log(`Number of schedule items found: ${scheduleItems.length}`);
        if (scheduleItems.length === 0) {
            const pageContent = await page.content();
            console.log('Page content after search:', pageContent);
        }
        expect(scheduleItems.length).toBeGreaterThan(0);

        // Car Plate
        const carPlateDropdown = await page.locator("[data-testid='search-carplate']");
        console.log('Car plate dropdown locator:', carPlateDropdown);
        await carPlateDropdown.selectOption({ label: 'SFD3735F' });
        console.log('Selected "SFD3735F" from car plate dropdown');
        await page.waitForTimeout(2000);

        const scheduleItems1 = await page.$$('.schedule-item');
        console.log(`Number of schedule items found: ${scheduleItems1.length}`);
        if (scheduleItems1.length === 0) {
            const pageContent = await page.content();
            console.log('Page content after search:', pageContent);
        }
        expect(scheduleItems1.length).toBeGreaterThan(0);

        // Start Time
        const startTimeLocator = await page.locator("[data-testid='search-start-time']");
        console.log('Start time locator:', startTimeLocator);
        await startTimeLocator.fill('8:30');
        console.log('Filled search input with "8.30"');
        await page.waitForTimeout(2000);

        const scheduleItems2 = await page.$$('.schedule-item');
        console.log(`Number of schedule items found: ${scheduleItems2.length}`);
        if (scheduleItems2.length === 0) {
            const pageContent = await page.content();
            console.log('Page content after search:', pageContent);
        }
        expect(scheduleItems2.length).toBeGreaterThan(0);

        // End Time
        const endTimeLocator = await page.locator("[data-testid='search-end-time']");
        console.log('End time locator:', endTimeLocator);
        await endTimeLocator.fill('5.30');
        console.log('Filled search input with "5.30"');
        await page.waitForTimeout(2000);

        const scheduleItems3 = await page.$$('.schedule-item');
        console.log(`Number of schedule items found: ${scheduleItems3.length}`);
        if (scheduleItems3.length === 0) {
            const pageContent = await page.content();
            console.log('Page content after search:', pageContent);
        }
        expect(scheduleItems3.length).toBeGreaterThan(0);
    });
});
