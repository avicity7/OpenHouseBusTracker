import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Delete Bus', () => {
    test('Remove Bus from Bus Dashboard', async ({ page }) => {
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

        // Delete Button
        const deleteButtonSelector = '.bus-item button[data-testid="delete-bus-button"]';
        const busItems = await page.$$(`${deleteButtonSelector}`);
        if (busItems.length > 0) {
            console.log(`Deleting ${busItems.length} bus(s)...`);
            for (const scheduleItem of busItems) {
                for (const busItem of busItems) {
                    await busItem.click();
                    console.log('Bus deleted.');
            }
        }
        } else {
            console.log('No schedules found to delete.');
        }
    });
});
