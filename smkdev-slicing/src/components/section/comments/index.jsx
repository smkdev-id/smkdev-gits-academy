import Header from "../ui/Header";

const Comments = () => {
  // example datas
  const comment = [
    {
      name: "Hasban Fardani",
      position: "Siswa SMKN 11 Bandung",
      body: "Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!",
    },
    {
      name: "Picerld",
      position: "Siswa RPL",
      body: "Keren banget rasanya SMKDEV ini!!",
    },
    {
      name: "Rafi Cahyadi",
      position: "Siswa SMKN 11 Bandung",
      body: "Kerennn!!!!!!",
    },
  ];

  return (
    // example
    <div className="py-20">
      <Header
        title="Apa Kata Mereka?"
        description="Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah kamu ingin seperti mereka?"
      />
      <div className="flex flex-wrap items-center">
        {comment?.map((comment, i) => (
          <div key={i} className="w-full chat chat-start md:w-1/2 lg:w-1/3">
            <div className="w-full h-40 max-w-xs overflow-y-auto bg-gray-900 chat-bubble">
              <h3 className="text-lg text-white">{comment?.body}</h3>
            </div>
            <div className="mt-2 font-semibold chat-footer">
              {comment?.name} <br />
              <span className="text-sm font-normal">{comment.position}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Comments;
