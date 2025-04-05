//hooks
import { useState } from 'react'

//components
import Loading from './Loading'
import toast from 'react-hot-toast'

//interfaces
import { IUpdateProductRequest } from '@interfaces/product'

//redux
import { useUpdateProductMutation } from '@redux/services/product'

interface IProps {
  image_url: string
  product: IUpdateProductRequest
}

const ProductModalUpdate = (props: IProps) => {
  const { image_url, product } = props
  const [form, setForm] = useState<IUpdateProductRequest>({ ...product, image: null })
  const [UpdateProduct, { isLoading }] = useUpdateProductMutation()

  const handleChangeForm = (name: string, value: string) => {
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleUpdateProduct = async () => {
    const formData = new FormData()
    formData.append('id', product.id)
    formData.append('name', form.name)
    formData.append('description', form.description)
    formData.append('image', form.image)
    formData.append('price', form.price.toString())

    try {
      const result = await UpdateProduct(formData).unwrap()

      if (result) {
        const modal = document.getElementById(`update_product_modal_${product.id}`) as HTMLDialogElement
        modal.close()
        toast.success('Update product successfully.')
      }
    } catch (e: any) {
      toast.error('Something went wrong.')
    }
  }

  return (
    <dialog id={`update_product_modal_${product.id}`} className="modal">
      <div className="modal-box">
        <h3 className="text-lg font-bold">Add Product</h3>
        <p className="py-4">Please full fill information</p>
        <form className="flex flex-col gap-4">
          <fieldset className="fieldset">
            <legend className="fieldset-legend">Image</legend>
            {image_url && <img src={image_url} alt="" className="w-full h-[100px] object-cover" />}
            <input
              type="file"
              id="image"
              name="image"
              placeholder="Product Image"
              className="w-full file-input"
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
            <button onClick={handleUpdateProduct} className="btn btn-success">
              {isLoading ? <Loading /> : 'Update'}
            </button>
            <button className="btn">Close</button>
          </form>
        </div>
      </div>
    </dialog>
  )
}

export default ProductModalUpdate
