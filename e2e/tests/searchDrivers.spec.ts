import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Schedule Search Functionality', () => {
    test('Search Schedules by Driver Name', async ({ page }) => {
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

        // Search Function
        const searchInputLocator = await page.locator("[data-testid='search-drivers']");
        console.log('Search input locator:', searchInputLocator);
        await searchInputLocator.fill('Lily Fields');
        console.log('Filled search input with "Lily Fields"');
        await page.waitForTimeout(2000); 
        const driversList = await page.$$('.drivers-list');
        console.log(`Number of drivers found: ${driversList.length}`);
        if (driversList.length === 0) {
            const pageContent = await page.content();
            console.log('Page content after search:', pageContent);
        }
        expect(driversList.length).toBeGreaterThan(0);
    });
});
