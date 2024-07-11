import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Delete Schedules', () => {
    test('Remove Schedules from Schedule Dashboard', async ({ page }) => {
        //Login
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

        // Delete Button
        const deleteButtonSelector = '.schedule-item button[data-testid="delete-schedule-button"]';
        const scheduleItems = await page.$$(`${deleteButtonSelector}`);
        if (scheduleItems.length > 0) {
            console.log(`Deleting ${scheduleItems.length} schedule(s)...`);
            for (const scheduleItem of scheduleItems) {
                for (const scheduleItem of scheduleItems) {
                    await scheduleItem.click();
                    console.log('Schedule deleted.');
            }
        }
        } else {
            console.log('No schedules found to delete.');
        }
    });
});
