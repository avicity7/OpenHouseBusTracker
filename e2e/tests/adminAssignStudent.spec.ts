import { test, expect } from '@playwright/test';

test.describe('Admin - Bus Assignment', () => {
  test('assigns a student to a different bus', async ({ page }) => {
    console.log('Navigating to admin login page...');
    // Login as admin
    await page.goto('http://localhost:5173/admin-login');
    await page.fill('#username', 'admin');
    await page.fill('#password', 'password');
    await page.click('#login-button');

    console.log('Logged in as admin, navigating to bus assignment page...');
    await page.goto('http://localhost:5173/bus-assignment');

    console.log('Selecting a student and assigning to a different bus...');
    // Assign a student to a different bus
    await page.selectOption('#student-select', 'student-id'); 
    await page.selectOption('#bus-select', 'new-bus-id'); 
    await page.click('#assign-bus-button');

    console.log('Waiting for assignment confirmation...');
    const assignmentNotification = await page.locator('.notification').textContent();
    console.log('Assignment Notification:', assignmentNotification);

    expect(assignmentNotification).toBe('Student assigned to new bus!!');
  });
});
