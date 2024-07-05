import { test, expect, devices } from '@playwright/test';

import { randomString } from './utils';

const username = `testuser-${randomString()}`;
const password = 'password';

test.describe('new user sign up', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/app/signup');
    });

    test('sign up with credentials', async ({ page }) => {
        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', password);
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page).toHaveURL('/app');
    });

    test('sign up with same credentials', async ({ page }) => {
        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', 'password');
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
    });
});
