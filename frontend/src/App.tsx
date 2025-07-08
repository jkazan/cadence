import React, { useState } from 'react';

function App() {
  const [artist, setArtist] = useState('');
  const [title, setTitle] = useState('');

const [songInfo, setSongInfo] = useState({
  tempo: '',
  danceability: '',
  acousticness: '',
  keyOf: '',
  openKey: '',
});


  const [loading, setLoading] = useState(false);


  const handleSearch = () => {
    if (!artist || !title) return;

    setLoading(true);

    fetch(`http://localhost:8080/api/tempo?artist=${encodeURIComponent(artist)}&title=${encodeURIComponent(title)}`)
      .then((res) => res.json())
      .then((data) => {
        setSongInfo({
            tempo: data.tempo,
            danceability: data.danceability,
            acousticness: data.acousticness,
            keyOf: data.key_of,
            openKey: data.open_key,
          }
        )
      })
      .catch(console.error)
      .finally(() => setLoading(false));
  };

  return (
    <div className="p-6 max-w-md mx-auto">
      <h1 className="text-xl font-bold mb-4">Tempo Finder</h1>

      <input
        type="text"
        placeholder="Artist"
        value={artist}
        onChange={(e) => setArtist(e.target.value)}
        className="border p-2 w-full mb-2 rounded"
      />
      <input
        type="text"
        placeholder="Title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        className="border p-2 w-full mb-4 rounded"
      />

      <button
        onClick={handleSearch}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        disabled={loading}
      >
        {loading ? 'Searching...' : 'Find Tempo'}
      </button>

      <div className="mt-6 text-lg">
        {songInfo && <p>Tempo (BPM): <strong>{songInfo.tempo}</strong></p>}
        {songInfo && <p>Danceabilityt (0-100): <strong>{songInfo.danceability}</strong></p>}
        {songInfo && <p>Acousticness (0-100): <strong>{songInfo.acousticness}</strong></p>}
        {songInfo && <p>Key of: <strong>{songInfo.keyOf}</strong></p>}
        {songInfo && <p>Open Key: <strong>{songInfo.openKey}</strong></p>}
      </div>
    </div>
  );
}

export default App;
