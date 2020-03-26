// Import `shallowMount` from Vue Test Utils and the component being tested
import { shallowMount } from '@vue/test-utils'
import FarmInfo from '../components/home/FarmInfo.vue'

// Mount the component
const wrapper = shallowMount(FarmInfo)

// Here are some Jest tests, though you can
// use any test runner/assertion library combo you prefer
describe('FarmInfo', () => {
  // Inspect the raw component options
  it('has a created hook', () => {
    expect(typeof FarmInfo.created).toBe('function')
  })
})