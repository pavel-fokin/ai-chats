import { test, expect } from '@playwright/test';

test.describe('signup page', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/app/signup');
    });

    test('load', async ({ page }) => {
        await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
        await expect(page.getByRole('button', { name: 'Create an account' })).toBeVisible();
        await expect(page.getByRole('link', { name: 'Log in' })).toBeVisible();
    });

    test('access login page', async ({ page }) => {
        await page.getByRole('link', { name: 'Log in' }).click();
        await expect(page.getByRole('heading', { name: 'Log in' })).toBeVisible();
    });

    test('sign up with empty credentials', async ({ page }) => {
        await page.getByRole('button', { name: 'Create an account' }).click();
        await expect(page.getByText('Username is required')).toBeVisible();
        await expect(page.getByText('Password must be at least 6 characters')).toBeVisible();
    });
});

