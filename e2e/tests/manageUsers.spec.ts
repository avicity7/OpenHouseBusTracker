import { test, expect } from '@playwright/test';
import SignInPage from '../models/SignInPage';

test.describe('Authorisation of Accounts', () => {
    test('Ensuring authorisation of accounts', async ({ page }) => {
        //Login
        let signInPage = new SignInPage(page);
        await signInPage.visit();
        const email = 'sarahaxl7@gmail.com'; 
        const password = '123'; 
        console.log('Attempting to sign in with provided credentials');
        await signInPage.signIn(email, password);
        const currentURL = page.url();
        console.log('Current URL after login:', currentURL);
        if (!currentURL.includes('/admin/users')) {
            console.log('Navigating manually to /admin/users');
            await page.goto('http://open-house-bus-tracker.vercel.app/admin/users');
        }
        const finalURL = page.url();
        expect(finalURL).toContain('/admin/users');
        await page.screenshot({ path: 'admin-users-page.png' });

        const roleDropdown = await page.locator('[data-testid="update-role-lalibiluo@gmail.com"]');
        await roleDropdown.click(); 
        await roleDropdown.selectOption({ label: 'admin' });
    });
});
