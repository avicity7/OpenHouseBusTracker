import { test, expect } from '@playwright/test';
import StudentBusTrackPage from '../models/BusTrackPage'; 

test.describe('Student Bus Tracking', () => {

  test('should display bus arrival status and handle WebSocket updates', async ({ page }) => {
    const studentBusTrackPage = new StudentBusTrackPage(page);
    await studentBusTrackPage.visitSignIn();
    const email = 'user@gmail.com';
    const password = '1'; 
    await studentBusTrackPage.signIn(email, password);
    await studentBusTrackPage.visitBusTrack();
    await studentBusTrackPage.simulateBusArrival();
    await studentBusTrackPage.checkBusArrivalStatus('Bus Arrived!!'); 
    await studentBusTrackPage.checkStopDetailsVisible();
  });

});
