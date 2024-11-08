import { test, expect, devices, BrowserContext, Page } from '@playwright/test';

import { randomString } from './utils';


test.describe('user sign up', () => {
  let context: BrowserContext;
  let page: Page;

  const username = `testuser-${randomString()}`;
  const password = 'password';

  test.beforeAll(async ({ browser }) => {
    context = await browser.newContext();
    page = await context.newPage();

    await page.goto('/app/signup');
    await page.fill('input[name="username"]', username);
    await page.fill('input[name="password"]', password);
    await page.getByRole('button', { name: 'Create an account' }).click();
    await expect(page).toHaveURL('/app/new-chat');

    // Sign out to allow for login tests later.
    await page.getByLabel('Sign Out').click();
  });

  // Close the context after all tests.
  test.afterAll(async () => {
    await context.close();
  });

  test.describe('user login tests', () => {
    test.beforeEach(async ({ page }) => {
      await page.goto('/app/login');
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

    test('login with credentials', async ({ page }) => {
      await page.fill('input[name="username"]', username);
      await page.fill('input[name="password"]', password);
      await page.getByRole('button', { name: 'Log in' }).click();
      await expect(page).toHaveURL('/app/new-chat');
    });

    test('navigate to ollama settings', async ({ page }) => {
      await page.fill('input[name="username"]', username);
      await page.fill('input[name="password"]', password);
      await page.getByRole('button', { name: 'Log in' }).click();

      await page.getByRole('link', { name: 'Ollama Settings' }).click();
      await expect(page).toHaveURL('/app/settings');
    });
  });
});

test.describe('user is chatting', () => {
  const username = `testuser-${randomString()}`;
  const password = 'password';

  test.beforeEach(async ({ page }) => {
    await page.goto('/app/signup');
    await page.fill('input[name="username"]', username);
    await page.fill('input[name="password"]', password);
    await page.getByRole('button', { name: 'Create an account' }).click();
    await expect(page).toHaveURL('/app/new-chat');
  });

  test('start new chat', async ({ page }) => {
    await page.goto('/app/new-chat');
    await page.fill('textarea', 'Hello, world!');
    await page
      .getByRole('button', { name: 'Send a message', exact: true })
      .click();
    await expect(page.getByText('Hello, world!')).toBeVisible();
    await expect(page.getByText(/^Model/, { exact: true })).toBeVisible();
  });

  test('start new chat and delete it', async ({ page }) => {
    await expect(page).toHaveURL('/app/new-chat');
    await expect(
      page.getByRole('heading', { name: 'New chat' })
    ).toBeVisible();

    // Send a message.
    await page.fill('textarea', 'Hello, world!');
    await page
      .getByRole('button', { name: 'Send a message', exact: true })
      .click();
    await expect(page.getByText('Hello, world!')).toBeVisible();
    await expect(page.getByText(/^Model/, { exact: true })).toBeVisible();

    // Delete the chat.
    await page.getByLabel('Chat menu button').click();
    await page.getByRole('menuitem', { name: 'Delete chat' }).click();
    await expect(page.getByText('Delete chat?')).toBeVisible();
    await page.getByRole('button', { name: 'Delete chat' }).click();

    await expect(page).toHaveURL('/app/new-chat');
  });
});
