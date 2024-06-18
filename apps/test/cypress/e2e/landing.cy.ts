/// <reference types="cypress" />

const BASE_URL = process.env.AIBOTS_CYPRESS_BASE_URL || 'http://localhost:8080'

const sizes = ['iphone-6', 'macbook-13']

describe('Landing Page', () => {
  sizes.forEach((size) => {
    it(`[${size}] it should have title`, () => {
      cy.viewport(size as Cypress.ViewportPreset)
      cy.visit(BASE_URL)
      cy.contains('Create and Manage Your AI Chats')
    })

    it(`[${size}] it should open 'Log in' page`, () => {
      cy.viewport(size as Cypress.ViewportPreset)
      cy.visit(BASE_URL)
      cy.contains('Create and Manage Your AI Chats')
      cy.get('a').contains('Log in').click()
      cy.get('h2').contains('Log in')
    })

    it(`[${size}] it should open 'Sign up' page`, () => {
      cy.viewport(size as Cypress.ViewportPreset)
      cy.visit(BASE_URL)
      cy.contains('Create and Manage Your AI Chats')
      cy.get('a').contains('Sign up').click()
      cy.get('h2').contains('Sign up')
    })
  })
})
