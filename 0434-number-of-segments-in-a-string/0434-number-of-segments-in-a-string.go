func countSegments(s string) int {
    segments := strings.Fields(s)
    return len(segments)
}