import { test, expect } from '@playwright/test';

test.describe('login page', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/app/login');
    });

    test('load', async ({ page }) => {
        await expect(page.getByRole('heading', { name: 'Log in' })).toBeVisible();
        await expect(page.getByRole('button', { name: 'Log in' })).toBeVisible();
        await expect(page.getByRole('link', { name: 'Create one' })).toBeVisible();
    });

    test('access signup page', async ({ page }) => {
        await page.getByRole('link', { name: 'Create one' }).click();
        await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
    });

    test('log in with empty credentials', async ({ page }) => {
        await page.getByRole('button', { name: 'Log in' }).click();
        await expect(page.getByText('Username is required')).toBeVisible();
        await expect(page.getByText('Password must be at least 6 characters')).toBeVisible();
    });
});

