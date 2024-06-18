import { test, expect } from '@playwright/test';

test.describe('landing page', () => {
    test.beforeEach(async ({ page }) => {
        await page.goto('/');
    });

    test('load', async ({ page }) => {
        await expect(page).toHaveTitle('AI Chats');
        await expect(page.getByRole('heading', { name: 'Create and Manage Your AI Chats' })).toBeVisible();
        await expect(page.getByRole('link', { name: 'Log in' })).toBeVisible();
        await expect(page.getByRole('link', { name: 'Sign up' })).toBeVisible();
    });

    test('access log in', async ({ page }) => {
        await page.getByRole('link', { name: 'Log in' }).click();
        await expect(page.getByRole('heading', { name: 'Log in' })).toBeVisible();
    });

    test('access sign up', async ({ page }) => {
        await page.getByRole('link', { name: 'Sign up' }).click();
        await expect(page.getByRole('heading', { name: 'Sign up' })).toBeVisible();
    });
});