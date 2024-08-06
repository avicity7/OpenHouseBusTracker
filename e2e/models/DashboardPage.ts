import { Page } from '@playwright/test';

export default class DashboardPage {
  page: Page;
  
  constructor(page: Page) {
    this.page = page;
  }

  get studentWelcomeMessage() {
    return this.page.locator('h1:has-text("Welcome, Student!")');
  }

  get currentShiftsHeader() {
    return this.page.locator('h1:has-text("Current Shifts")');
  }

  get upcomingShiftsHeader() {
    return this.page.locator('h1:has-text("Upcoming Shifts")');
  }
}