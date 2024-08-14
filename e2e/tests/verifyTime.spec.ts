import { test, expect } from '@playwright/test';

test.describe('Admin - Real-Time Updates', () => {
  test('verifies real-time updates for student absences', async ({ page }) => {
    console.log('Navigating to admin login page...');
    // Login as admin
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Logged in as admin, navigating to the real-time updates page');
    // Navigate to the real-time updates page
    await page.goto('http://localhost:5173/real-time-updates');

    console.log('Logging student absence to trigger real-time update');
    // marking a student as absent
    await page.goto('http://localhost:5173/absence-management');
    await page.selectOption('#student-select', 'student-id');
    await page.fill('#absence-reason', 'Sick');
    await page.click('#mark-absent-button');

    console.log('Waiting for the real-time update to be reflected');
    // Wait real-time update to appear on the page
    const realTimeUpdate = await page.locator('.real-time-update').textContent();
    console.log('Real-Time Update:', realTimeUpdate);

    // Verify that the real-time update reflects the student absence
    expect(realTimeUpdate).toContain('Student marked as absent.');
  });
});
