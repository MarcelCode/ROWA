// Import `shallowMount` from Vue Test Utils and the component being tested
import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import FarmInfo from '../components/home/FarmInfo.vue'
import FarmTransition from '../components/main/FarmTransition.vue'
import InfoBoxPlants from '../components/home/InfoBoxPlants.vue'
import vuetify from "vuetify"
import axios from 'axios';


jest.mock('axios');
axios.get.mockResolvedValue({})
// Mount the component
//const wrapper = shallowMount(FarmInfo)
// Here are some Jest tests, though you can
// use any test runner/assertion library combo you prefer
describe('FarmInfo',  () => {
  let wrapper;
  beforeEach(() => {
    const localVue =  createLocalVue()
    localVue.use(vuetify)
    wrapper = shallowMount(FarmInfo, {localVue})    
  })

  
  
  // Inspect the raw component options
  it('has a created hook',()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })
  it('contains the right components', ()=>{
    expect(wrapper.contains(FarmTransition)).toBe(true)
    expect(wrapper.contains(InfoBoxPlants)).toBe(true)
  })
  it('displays the correct amount of harvestable plants', async() =>{
    //wrapper.vm.harvestable = null
    wrapper = shallowMount(FarmInfo)
    console.log(wrapper.vm.harvestable)
    const data = {
      data: 1
    }
    axios.get.mockResolvedValue(data)
    await wrapper.vm.$nextTick()
    //axios.get.mockImplementationOnce(() => Promise.resolve(data));
    console.log(wrapper.vm.harvestable)
    //expect().toBe(5)
  })

  it('displays the correct amount of plantable plants', () =>{
    axios.get.mockResolvedValue({data:'4'})
    //expect(FarmInfo.harvestable).toBe(4)
  })
})