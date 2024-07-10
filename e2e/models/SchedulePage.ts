import { type Page, type Locator, expect } from "@playwright/test";

export default class DriversPage {
    readonly page: Page;
    readonly url: string = 'http://localhost:5173/admin/driver';
    readonly searchInput: Locator;
    readonly searchButton: Locator;
    readonly driverList: Locator;

    constructor(page: Page) {
        this.page = page;
        this.searchInput = page.getByTestId("drivers-search-input");
        this.searchButton = page.getByRole('button', { name: "Search" });
        this.driverList = page.getByTestId("drivers-list");
    }

    async visit() {
        await this.page.goto(this.url);
    }

    async searchForDriversOnDuty() {
        await this.searchInput.fill('on duty');
        await this.searchButton.click();
    }

    async getDriversOnDuty() {
        return this.driverList.locator('.driver').allTextContents();
    }
}
