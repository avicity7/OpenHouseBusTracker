import { Page, Locator, expect } from '@playwright/test';

export default class StudentSchedulePage {
    readonly page: Page;
    readonly scheduleUrl: string = '/schedules'; 
    readonly currentSchedulesContainer: Locator;
    readonly upcomingSchedulesContainer: Locator;
    readonly noCurrentSchedulesMessage: Locator;
    readonly noUpcomingSchedulesMessage: Locator;

    constructor(page: Page) {
        this.page = page;
        this.currentSchedulesContainer = page.locator("[data-testid='current-schedules']");
        this.upcomingSchedulesContainer = page.locator("[data-testid='upcoming-schedules']");
        this.noCurrentSchedulesMessage = page.locator("[data-testid='no-current-schedules']");
        this.noUpcomingSchedulesMessage = page.locator("[data-testid='no-upcoming-schedules']");
    }
    async visit() {
        await this.page.goto(this.scheduleUrl);
    }
    async verifyCurrentSchedulesVisible() {
        await expect(this.currentSchedulesContainer).toBeVisible();
    }
    async verifyUpcomingSchedulesVisible() {
        await expect(this.upcomingSchedulesContainer).toBeVisible();
    }
    async verifyNoCurrentSchedulesMessage() {
        await expect(this.noCurrentSchedulesMessage).toBeVisible();
    }
    async verifyNoUpcomingSchedulesMessage() {
        await expect(this.noUpcomingSchedulesMessage).toBeVisible();
    }
}
