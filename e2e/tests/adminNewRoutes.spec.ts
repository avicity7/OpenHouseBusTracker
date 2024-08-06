import { test, expect } from '@playwright/test';

test.describe('Admin', () => {
  test('View routes, route steps, and add a new route', async ({ page }) => {
    // Admin login
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    // Navigate to routes page
    await page.goto('http://localhost:5173/admin/bus-routes');

    // View routes
    console.log('Routes:');
    const routeNames = await page.locator('table:has-text("Route Name")').allTextContents();
    console.log(routeNames.join('\n'));

    // View route steps
    console.log('Route Steps:');
    const routeSteps = await page.locator('table:has-text("Stop Name")').allTextContents();
    console.log(routeSteps.join('\n'));

    // Add new route
    await page.goto('http://localhost:5173/admin/bus-routes/create-route');
    await page.fill('input[name="RouteName"]', 'New Route');
    await page.fill('input[type="color"]', '#ff0000');
    await page.click('button[type="submit"]');

    // Verify route added
    const updatedRouteNames = await page.locator('table:has-text("Route Name")').allTextContents();
    expect(updatedRouteNames).toContain('New Route');
  });
});
