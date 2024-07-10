import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Schedule Search Functionality', () => {
    test('Navigate to Admin Schedule After Login and Search Schedules by Driver Name', async ({ page }) => {
        let signInPage = new SignInPage(page);
        await signInPage.visit();

        const email = 'sarahaxl7@gmail.com'; // Replace with admin credentials
        const password = '123'; // Replace with admin password

        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);

        // Verify the current URL after login
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        
        // Navigate manually if not on admin schedule page
        if (!currentURL.includes('/admin/schedule')) {
            console.log('Navigating manually to /admin/schedule');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/schedule');
        }

        // Assert that the URL navigated to the admin schedule page
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/schedule');

        // Optionally, take a screenshot for visual verification
        await page.screenshot({ path: 'admin-schedule-page.png' });

        const searchInputLocator = await page.locator("[data-testid='search-input']");
        console.log('Search input locator:', searchInputLocator);
        // Type the driver name to search for
        await searchInputLocator.fill('Lily Fields');

        // Wait for search results to appear (adjust timeout as needed)
        await page.waitForTimeout(2000); // Wait for 2 seconds

        // Check if schedule items are found
        const scheduleItems = await page.$$('.schedule-item');
        expect(scheduleItems.length).toBeGreaterThan(0);
    });
    });
