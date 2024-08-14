import { test, expect } from '@playwright/test';
import StudentSchedulePage from '../models/StudentSchedulePage';
import SignInPage from '../models/SignInPage';

test.describe('Student Schedule Viewing', () => {

  test('should display current and upcoming schedules after login', async ({ page }) => {
    const signInPage = new SignInPage(page);
    const studentSchedulePage = new StudentSchedulePage(page);

    //Sign in as student
    await signInPage.visit();
    const email = 'user@gmail.com'; 
    const password = '1'; 
    await signInPage.signIn(email, password);

    //Navigate schedule page
    await studentSchedulePage.visit();

    //Verify current and upcoming schedules displays
    await studentSchedulePage.verifyCurrentSchedulesVisible();
    await studentSchedulePage.verifyUpcomingSchedulesVisible();
    await studentSchedulePage.verifyNoCurrentSchedulesMessage(); 
    await studentSchedulePage.verifyNoUpcomingSchedulesMessage(); 
  });

});
