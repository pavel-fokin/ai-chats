import { test, expect, devices } from '@playwright/test';

import { randomString } from './utils';

const username = `testuser-${randomString()}`;
const password = 'password';

test.describe('user sign up', () => {
  let context;
  let page;

  test.beforeAll(async ({ browser }) => {
    context = await browser.newContext();
    page = await context.newPage();

    await page.goto('/app/signup');
    await page.fill('input[name="username"]', username);
    await page.fill('input[name="password"]', password);
    await page.getByRole('button', { name: 'Create an account' }).click();
    await expect(page).toHaveURL('/app/new-chat');

    // Sign out to allow for login tests later
    await page.getByLabel('Sign Out').click();
  });

  // Close the context after all tests
  test.afterAll(async () => {
    await context.close();
  });

  test.describe('user login tests', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/app/login');
    });

    test('login with credentials', async ({ page }) => {
      await page.fill('input[name="username"]', username);
      await page.fill('input[name="password"]', password);
      await page.getByRole('button', { name: 'Log in' }).click();
      await expect(page).toHaveURL('/app/new-chat');
    });

    test('start new chat after login', async ({ page }) => {
      await page.fill('input[name="username"]', username);
      await page.fill('input[name="password"]', password);
      await page.getByRole('button', { name: 'Log in' }).click();

      await expect(page).toHaveURL('/app/new-chat');
      await expect(
        page.getByRole('heading', { name: 'New chat' })
      ).toBeVisible();
      await page.fill('textarea', 'Hello, world!');
      await page.getByRole('button', { name: /Send a message/, exact: true }).click();
      await expect(page.getByText(/Hello, world!/)).toBeVisible();
      await expect(page.getByText(/Model/)).toBeVisible();
    });
  });

  test('attempt to sign up with existing credentials', async ({ page }) => {
    await page.goto('/app/signup');
    await page.fill('input[name="username"]', username);
    await page.fill('input[name="password"]', password);
    await page.getByRole('button', { name: 'Create an account' }).click();
    await expect(
      page.getByText('That username is already taken. Try another one.')
    ).toBeVisible();
  });
});
