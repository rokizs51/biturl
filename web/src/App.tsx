import React, { useState } from 'react';
import { Link2, Loader2, Copy, ExternalLink, Check } from 'lucide-react';

interface ShortenedUrl {
  originalUrl: string;
  shortUrl: string;
}

// Updated interface to match your backend response
interface ApiResponse {
  short_url: string;
  long_url: string;
  expires: string;
}

function App() {
  const [url, setUrl] = useState('');
  const [customSlug, setCustomSlug] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [result, setResult] = useState<ShortenedUrl | null>(null);
  const [error, setError] = useState('');
  const [copied, setCopied] = useState(false);
  const [urlError, setUrlError] = useState('');

  // URL validation function
  const isValidUrl = (urlString: string): boolean => {
    try {
      const url = new URL(urlString);
      return url.protocol === 'http:' || url.protocol === 'https:';
    } catch (err) {
      return false;
    }
  };

  // Handle URL input change with validation
  const handleUrlChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setUrl(value);
    
    if (value && !isValidUrl(value)) {
      setUrlError('Please enter a valid URL starting with http:// or https://');
    } else {
      setUrlError('');
    }
  };

  // Add new state for custom slug validation
  const [customSlugError, setCustomSlugError] = useState('');
  
  // Add custom slug validation function
  const isValidCustomSlug = (slug: string): boolean => {
    const slugRegex = /^[a-zA-Z0-9-]{3,15}$/;
    return slugRegex.test(slug);
  };
  
  // Add custom slug change handler with validation
  const handleCustomSlugChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setCustomSlug(value);
    
    if (value && !isValidCustomSlug(value)) {
      setCustomSlugError('Custom slug must be 3-15 characters long and contain only letters, numbers, and hyphens');
    } else {
      setCustomSlugError('');
    }
  };
  
  // Update handleSubmit to include custom slug validation
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!isValidUrl(url)) {
      setUrlError('Please enter a valid URL starting with http:// or https://');
      return;
    }
  
    if (customSlug && !isValidCustomSlug(customSlug)) {
      setCustomSlugError('Custom slug must be 3-15 characters long and contain only letters, numbers, and hyphens');
      return;
    }
    
    setIsLoading(true);
    setError('');
    setResult(null);

    try {
      // Get API URL from environment variables
      const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';
      console.log(import.meta.env.VITE_API_URL)
      // Make the actual API call to your backend
      const response = await fetch(`${apiUrl}api/shorten`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ long_url: url, custom_slug : customSlug }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to shorten URL');
      }

      const data: ApiResponse = await response.json();
      
      setResult({
        originalUrl: data.long_url,
        shortUrl: import.meta.env.VITE_API_URL +  data.short_url
      });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to shorten URL. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const copyToClipboard = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy:', err);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-violet-50 via-blue-50 to-emerald-50 p-4">
      <div className="max-w-2xl mx-auto pt-16 px-4">
        <div className="text-center mb-12">
          <div className="inline-block p-4 bg-white rounded-2xl shadow-md mb-6 transform hover:rotate-12 transition-transform duration-300">
            <Link2 className="w-10 h-10 text-violet-600" />
          </div>
          <h1 className="text-5xl font-bold bg-gradient-to-r from-violet-600 to-blue-600 text-transparent bg-clip-text mb-4">
            Bit URL 
          </h1>
          <p className="text-gray-600 text-lg">Transform long URLs into elegant, shareable links</p>
        </div>

        <div className="bg-white/80 backdrop-blur-lg rounded-3xl shadow-xl p-8 border border-white/20">
          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="url" className="block text-sm font-medium text-gray-700 mb-2">
                Enter your long URL
              </label>
              <input
                type="url"
                id="url"
                value={url}
                onChange={handleUrlChange}
                placeholder="https://example.com/very/long/url"
                className={`w-full px-6 py-4 bg-white border ${urlError ? 'border-red-300 focus:ring-red-500' : 'border-gray-200 focus:ring-violet-500'} rounded-2xl focus:ring-2 focus:border-transparent transition-all duration-200 outline-none text-gray-800 placeholder-gray-400`}
                required
              />
              {urlError && (
                <p className="mt-2 text-sm text-red-600">{urlError}</p>
              )}
            </div>
            
            {/* Modified section: Button and custom slug side by side */}
            <div className="flex space-x-4">
              <button
                type="submit"
                disabled={isLoading}
                className="w-1/2 bg-gradient-to-r from-violet-600 to-blue-600 text-white py-4 px-6 rounded-2xl hover:from-violet-700 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 font-medium text-lg shadow-lg shadow-violet-500/20"
              >
                {isLoading ? (
                  <span className="flex items-center justify-center">
                    <Loader2 className="w-5 h-5 animate-spin mr-2" />
                    Shortening...
                  </span>
                ) : (
                  'Shorten URL'
                )}
              </button>
              
              <div className="w-1/2">
                <input
                  type="text"
                  id="custom_slug"
                  value={customSlug}
                  onChange={handleCustomSlugChange}
                  placeholder="Custom slug (optional)"
                  className={`w-full px-6 py-4 bg-white border ${
                    customSlugError ? 'border-red-300 focus:ring-red-500' : 'border-gray-200 focus:ring-violet-500'
                  } rounded-2xl focus:ring-2 focus:border-transparent transition-all duration-200 outline-none text-gray-800 placeholder-gray-400`}
                />
                {customSlugError && (
                  <p className="mt-2 text-sm text-red-600">{customSlugError}</p>
                )}
              </div>
            </div>

          </form>

          {error && (
            <div className="mt-6 p-6 bg-red-50 text-red-700 rounded-2xl border border-red-100 animate-fade-in">
              {error}
            </div>
          )}

          {result && (
            <div className="mt-8 space-y-6 animate-fade-in">
              <div className="p-6 bg-violet-50 rounded-2xl border border-violet-100">
                <h3 className="text-xl font-semibold text-violet-900 mb-4">Your shortened URL is ready!</h3>
                <div className="space-y-4">
                  <div className="space-y-2">
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-violet-700">Original URL</span>
                      <a
                        href={result.originalUrl}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-violet-600 hover:text-violet-800 inline-flex items-center text-sm"
                      >
                        <ExternalLink className="w-4 h-4 ml-1" />
                      </a>
                    </div>
                    <p className="text-gray-600 break-all bg-white/50 rounded-xl p-3 text-sm">
                      {result.originalUrl}
                    </p>
                  </div>
                  <div className="space-y-2">
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-violet-700">Short URL</span>
                      <button
                        onClick={() => copyToClipboard(result.shortUrl)}
                        className="text-violet-600 hover:text-violet-800 inline-flex items-center gap-1 text-sm"
                      >
                        {copied ? (
                          <>
                            Copied! <Check className="w-4 h-4" />
                          </>
                        ) : (
                          <>
                            Copy <Copy className="w-4 h-4" />
                          </>
                        )}
                      </button>
                    </div>
                    <div className="bg-white rounded-xl p-3 flex items-center justify-between group relative">
                      <a
                        href={result.shortUrl}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-blue-600 hover:text-blue-800 break-all text-sm font-medium"
                      >
                        {result.shortUrl}
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>

        <footer className="mt-12 text-center text-gray-500 text-sm">
          Made with ❤️ for sharing links efficiently
        </footer>
      </div>
    </div>
  );
}

export default App;