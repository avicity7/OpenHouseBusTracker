import { test, expect } from '@playwright/test';
import searchByDriverName from '../models/SearchDriverName';

test.describe('Admin Schedule Search', () => {
  test('Search schedules by driver name', async ({ page }) => {
    const adminSchedulePage = new searchByDriverName(page);

    // Perform login before visiting the schedule page
    await adminSchedulePage.login('user@gmail.com', '1');

    await adminSchedulePage.visit();
    await adminSchedulePage.searchByDriverName('Lily Fields');
    const scheduleItemCount = await adminSchedulePage.getScheduleItemCount();
    expect(scheduleItemCount).toBeGreaterThan(0);
  });

  test('Search schedules with no results', async ({ page }) => {
    const adminSchedulePage = new searchByDriverName(page);

    // Perform login before visiting the schedule page
    await adminSchedulePage.login('admin@example.com', 'password');

    await adminSchedulePage.visit();
    await adminSchedulePage.searchByDriverName('Nonexistent Driver');
    const scheduleItemCount = await adminSchedulePage.getScheduleItemCount();
    expect(scheduleItemCount).toBe(0);
  });
});
