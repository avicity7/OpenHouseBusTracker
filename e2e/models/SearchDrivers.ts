import { Page, Locator } from "@playwright/test";

export default class SchedulePage {
    readonly page: Page;

    readonly url: string = '/admin/schedules';

    // Update to use data-testid for search input and schedule rows
    readonly searchInput: Locator;
    readonly scheduleRows: Locator;

    constructor(page: Page) {
        this.page = page;

        // Adjusted to use data-testid for search input and schedule rows
        this.searchInput = page.locator('[data-testid="search-input"]');
        this.scheduleRows = page.locator('[data-testid="search-results"]');
    }

    async visit() {
        await this.page.goto(this.url);
    }

    async searchByDriverName(driverName: string) {
        await this.searchInput.fill(driverName);
        await this.page.waitForTimeout(1000); // Adjust wait time if needed
    }

    async getFilteredScheduleCount() {
        const scheduleCount = await this.scheduleRows.count();
        return scheduleCount;
    }
}
