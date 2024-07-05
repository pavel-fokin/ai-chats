import { test, expect, devices } from '@playwright/test';

import { randomString } from './utils';

test.describe('new user sign up', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/app/signup');
    });

    const username = `testuser-${randomString()}`;

    test('sign up with credentials', async ({ page }) => {
        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', 'password');
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page).toHaveURL('/app');
        // await expect(page.getByText('AI Chats')).toBeVisible();
    });

    test('sign up with same credentials', async ({ page }) => {
        await page.fill('input[name="username"]', username);
        await page.fill('input[name="password"]', 'password');
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
    });
});
