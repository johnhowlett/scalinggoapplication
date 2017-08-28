package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// CloseableResponseWriter ist ein Interface das ein ResponsWriter und eine Close Function hat
type CloseableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (w gzipResponseWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func (w gzipResponseWriter) Close() {
	w.Writer.Close()
}

func (w gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (w closeableResponseWriter) Close() {}

//GetResponseWriter schaut ob gzip möglich ist
func GetResponseWriter(w http.ResponseWriter, req *http.Request) CloseableResponseWriter {
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gRW := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
		return gRW
	} else {
		return closeableResponseWriter{ResponseWriter: w}
	}
}

//GzipHandler dieser struct wird nur angelegt damit eine function erstellt werden kann
type GzipHandler struct{}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseWriter := GetResponseWriter(w, r)
	defer responseWriter.Close()

	http.DefaultServeMux.ServeHTTP(responseWriter, r)
}
