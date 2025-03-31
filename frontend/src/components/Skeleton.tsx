const Skeleton = () => {
  return (
    <div className="min-h-screen flex w-full flex-col gap-4 mt-20">
      <div className="skeleton h-32 w-full"></div>
      <div className="skeleton h-4 w-28"></div>
      <div className="skeleton h-4 w-full"></div>
      <div className="skeleton h-4 w-full"></div>
    </div>
  )
}

export default Skeleton
