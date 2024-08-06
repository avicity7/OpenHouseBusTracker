import { test, expect } from '@playwright/test';

test.describe('Admin Management', () => {

  test('should display routes and route steps', async ({ page }) => {
    console.log('Logging in as admin...');
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Navigating to routes page...');
    await page.goto('http://localhost:5173/admin/bus-routes');

    const routesLocator = page.locator('table tbody tr');
    const routesCount = await routesLocator.count();
    expect(routesCount).toBeGreaterThan(0); // Check routes are displayed

    for (let i = 0; i < routesCount; i++) {
      const routeRow = routesLocator.nth(i);
      const routeName = await routeRow.locator('td:first-child').textContent();
      console.log(`Route Name: ${routeName}`);
      
      const stepsLocator = page.locator(`text=${routeName} + table tbody tr`);
      const stepsCount = await stepsLocator.count();
      expect(stepsCount).toBeGreaterThan(0); // Check route steps are displayed
    }
  });

  test('should create and delete a route', async ({ page }) => {
    console.log('Logging in as admin...');
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Navigating to create route page...');
    await page.goto('http://localhost:5173/admin/bus-routes/create-route');

    console.log('Creating new route...');
    await page.fill('input[name="routeName"]', 'New Route');
    await page.fill('input[name="routeColor"]', '#ff0000'); // Set color
    await page.click('button[type="submit"]');

    console.log('Checking if new route appears...');
    await page.goto('http://localhost:5173/admin/bus-routes');
    const newRoute = page.locator('text=New Route');
    expect(await newRoute.isVisible()).toBe(true);

    console.log('Deleting the new route...');
    await newRoute.locator('button[title="Delete Route"]').click();
    await page.waitForTimeout(1000);

    console.log('Checking if route is deleted...');
    await page.goto('http://localhost:5173/admin/bus-routes');
    expect(await newRoute.isVisible()).toBe(false);
  });

  test('should create and delete a route step', async ({ page }) => {
    console.log('Logging in as admin...');
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Navigating to create route step page...');
    await page.goto('http://localhost:5173/admin/bus-routes/create-route-step/ExistingRoute'); // check my route for me ;-;

    console.log('Creating new route step...');
    await page.selectOption('select[name="stopName"]', 'New Stop');
    await page.click('button[type="submit"]');

    console.log('Checking if new route step appears...');
    await page.goto('http://localhost:5173/admin/bus-routes');
    const newStop = page.locator('text=New Stop');
    expect(await newStop.isVisible()).toBe(true);

    console.log('Deleting the new route step...');
    await newStop.locator('button[title="Delete Route Step"]').click();
    await page.waitForTimeout(1000); 

    console.log('Checking if route step is deleted...');
    await page.goto('http://localhost:5173/admin/bus-routes');
    expect(await newStop.isVisible()).toBe(false);
  });

  test('should view maps and building names', async ({ page }) => {
    console.log('Logging in as admin...');
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Navigating to maps and buildings page...');
    await page.goto('http://localhost:5173/admin/maps-and-buildings');

    const mapsLocator = page.locator('.map-container');
    const buildingsLocator = page.locator('.building-name');

    const mapsCount = await mapsLocator.count();
    const buildingsCount = await buildingsLocator.count();

    console.log(`Maps count: ${mapsCount}`);
    console.log(`Buildings count: ${buildingsCount}`);

    expect(mapsCount).toBeGreaterThan(0); // Check if maps are displayed
    expect(buildingsCount).toBeGreaterThan(0); // Check if buildings are displayed
  });

});
