module.exports = {
    'Testing frontend' : function (browser) {
      browser
        .url('http://localhost:8080')
        .waitForElementVisible('body')
        .assert.titleContains('frontend')
        .assert.visible('a[id=button]')
        .click('a[id=button]')
        .assert.urlEquals('http://localhost:8080/farming')
        .assert.containsText('p[data-v-2642cca2=""]', 'Go on a journey to become a famer.Harvest a plant now, for your daily dose of fresh greens.')
        .end();
    }
  };