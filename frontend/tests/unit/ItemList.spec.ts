import { mount } from '@vue/test-utils'
import ItemList from '@/components/ItemList.vue'

const wrapper = mount(ItemList)

describe('ItemList.vue', () => {
    it('renders item list', () => {
        console.log(wrapper.html())
        expect(wrapper.text().includes("No.")).toBe(true)
    })
})
