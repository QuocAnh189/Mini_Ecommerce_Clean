const ProductModalCreate = () => {
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
              placeholder="Product Image"
              className="w-full input input-bordered"
            />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Name</legend>
            <input type="text" placeholder="Enter Name" className="w-full input input-bordered" />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Description</legend>
            <input
              type="text"
              placeholder="Product Description"
              className="w-full input input-bordered"
            />
          </fieldset>

          <fieldset className="fieldset">
            <legend className="fieldset-legend">Price</legend>
            <input
              type="text"
              placeholder="Product Price"
              className="w-full input input-bordered"
            />
          </fieldset>
        </form>
        <div className="modal-action">
          <form method="dialog" className="flex gap-4">
            <button className="btn btn-success">Create</button>
            <button className="btn">Close</button>
          </form>
        </div>
      </div>
    </dialog>
  )
}

export default ProductModalCreate
