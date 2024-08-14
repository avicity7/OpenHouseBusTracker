import { test, expect } from '@playwright/test';

test.describe('Admin marks student', () => {
  test('marks a student as absent', async ({ page }) => {
    console.log('Navigating to admin login page...');
    // Login as admin
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Logged in as admin, navigating to absence management page...');
    await page.goto('http://localhost:5173/absence-management');
    console.log('Marking a student as absent...');

    // Mark a student as absent
    await page.selectOption('#student-select', 'student-id'); 
    await page.fill('#absence-reason', 'Sick'); 
    await page.click('#mark-absent-button'); 
    console.log('Waiting for notification...');

    // Verify the absence was marked
    const absenceNotification = await page.locator('.notification').textContent();
    console.log('Absence Notification:', absenceNotification);

    expect(absenceNotification).toBe('Student marked as absent.');
  });
});
