import { test, expect } from '@playwright/test';

test.describe('Admin Schedule Management', () => {

  test('Admin updates an existing schedule', async ({ page }) => {
    // Login as admin
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    // Navigate to the schedule management page
    await page.goto('http://localhost:5173/schedule-management');
    await page.click('text=Edit Schedule'); // help me adjust this ;-;
    await page.fill('#schedule-name', 'Updated Schedule Name');
    await page.fill('#schedule-details', 'Updated schedule details');
    await page.click('#save-button');

    // Verify schedule is updated
    const updatedNotification = await page.locator('.notification').textContent();
    expect(updatedNotification).toBe('Schedule updated successfully.');

    // Check updated schedule is displayed correctly
    const updatedSchedule = await page.locator('text=Updated Schedule Name').isVisible();
    expect(updatedSchedule).toBe(true);
  });

});
