import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Delete Event Helper', () => {
    test('Remove Event Helper from Event Helper Dashboard', async ({ page }) => {
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

        // Delete Button
        const deleteButtonSelector = '[data-testid="delete-helper-sarah"]';
        const helperItems = await page.$$(`${deleteButtonSelector}`);
        if (helperItems.length > 0) {
            console.log(`Deleting ${helperItems.length}`);
            for (const helperItem of helperItems) {
                for (const helperItem of helperItems) {
                    await helperItem.click();
                    console.log('Event Helper deleted.');
            }
        }
        } else {
            console.log('No event helper found to delete.');
        }
    });
});
