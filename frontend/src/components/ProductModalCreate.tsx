//hooks
import { useState } from 'react'

//components
import Loading from './Loading'
import toast from 'react-hot-toast'

//interfaces
import { ICreateProductRequest } from '@interfaces/product'

//redux
import { useCreateProductMutation } from '@redux/services/product'

const initForm: ICreateProductRequest = {
  name: '',
  description: '',
  image: null,
  price: 0,
}
const ProductModalCreate = () => {
  const [form, setForm] = useState<ICreateProductRequest>(initForm)

  const [CreateProduct, { isLoading }] = useCreateProductMutation()
  const handleChangeForm = (name: string, value: string) => {
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleCreateProduct = async () => {
    const formData = new FormData()
    formData.append('name', form.name)
    formData.append('description', form.description)
    formData.append('image', form.image)
    formData.append('price', form.price.toString())

    try {
      const result = await CreateProduct(formData).unwrap()

      if (result) {
        const modal = document.getElementById('create_product_modal') as HTMLDialogElement
        modal.close()
        setForm(initForm)
        toast.success('Create product successfully.')
      }
    } catch (e: any) {
      toast.error('Something went wrong.')
    }
  }

  return (
    <dialog id="create_product_modal" className="modal">
      <div className="modal-box">
        <h3 className="text-lg font-bold">Add Product</h3>
        <p className="py-4">Please full fill information</p>
        <form className="flex flex-col gap-4">
          <fieldset className="fieldset">
            <legend className="fieldset-legend">Image</legend>
            <input
              type="file"
              id="image"
              name="image"
              placeholder="Product Image"
              className="w-full input input-bordered"
              onChange={(e: any) => {
                handleChangeForm('image', e.target.files[0])
              }}
            />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Name</legend>
            <input
              id="name"
              name="name"
              value={form.name}
              type="text"
              placeholder="Enter Name"
              className="w-full input input-bordered"
              onChange={(e) => {
                handleChangeForm('name', e.target.value)
              }}
            />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Description</legend>
            <input
              id="description"
              name="description"
              value={form.description}
              type="text"
              placeholder="Product Description"
              className="w-full input input-bordered"
              onChange={(e) => {
                handleChangeForm('description', e.target.value)
              }}
            />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Price</legend>
            <input
              id="price"
              name="price"
              value={form.price}
              type="text"
              placeholder="Product Price"
              className="w-full input input-bordered"
              onChange={(e) => {
                handleChangeForm('price', e.target.value)
              }}
            />
          </fieldset>
        </form>
        <div className="modal-action">
          <form method="dialog" className="flex gap-4">
            <button onClick={handleCreateProduct} className="btn btn-success">
              {isLoading ? <Loading /> : 'Create'}
            </button>
            <button className="btn">Close</button>
          </form>
        </div>
      </div>
    </dialog>
  )
}

export default ProductModalCreate
