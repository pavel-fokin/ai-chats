import { test, expect, devices } from '@playwright/test';

import { randomString } from './utils';

test.describe('user sign up', () => {
    const username = `testuser-${randomString()}`;
    const password = 'password';

    test.beforeEach(async ({ page }) => {
        await page.goto('/app/signup');
    });

    test('sign up with credentials', async ({ page }) => {
        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', password);
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page).toHaveURL('/app/new-chat');

        // sign out
        await page.getByLabel('Sign Out').click();
        await page.goto('/app/signup');

        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', password);
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page.getByText('That username is already taken. Try another one.')).toBeVisible();
    });
});
